<?php

$file_name = "4.txt";
$file = fopen($file_name, 'r');

while ($line = fgets($file)) {
    $str = hex2bin(trim($line));
    $characters = str_split($str);

    $scores = [];

    // Score each character
    foreach ($characters as $character) {
        if (empty($scores[$character])) {
            $scores[$character] = 1;
        } else {
            $scores[$character] += 1;
        }
    }

    // Find the most common character
    foreach ($scores as $character => $score) {
        // Base case
        if (!isset($max)) {
            $max = $score;
            $max_char = $character;
            continue;
        }

        // Do we have a new max?
        if ($score > $max) {
            $max = $score;
            $max_char = $character;
        }
    }

    // From https://en.wikipedia.org/wiki/Letter_frequency#/media/File:English_letter_frequency_(alphabetic).svg
    $common_chars = array('a', 'e', 'h', 'i', 'n', 'o', 'r', 's', 't', ' ');
    $english_chars = array_map('chr', range(32, 126));

    $possible_results = [];
    foreach ($english_chars as $character) {
        $xor_result = $character ^ $max_char;
        if (in_array($xor_result, $common_chars)) {
            $possible_results[] = $character;
        }
    }

    foreach ($possible_results as $possible_result) {
        $result = [];

        foreach ($characters as $character) {
            $result[] = $character ^ $possible_result;
        }

        $decoded_result = join('', $result);

        if (preg_match('/^[A-Z](\w|\s|\')*[a-z]$/', $decoded_result)) {
            echo $possible_result . ": " . join('', $result) . "\n";
        }
    }

}
?>
