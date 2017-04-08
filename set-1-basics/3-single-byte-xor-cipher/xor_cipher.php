<?php

$hex_str = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736";
$str = hex2bin($hex_str);
$characters = str_split($str);

for ($i = 0; $i < 256; $i++) {
    $result = [];

    foreach ($characters as $character) {
        $result[] = $character ^ chr($i);
    }

    $result_str = join('', $result);
    if (
        preg_match('/^[A-Z].*[a-z]$/', $result_str) &&     // Does this make sense?
        strlen($result_str) == strlen($str)                 // Is it the same length?
    ) {
        echo chr($i) . ": " . $result_str . "\n";
    }
}

?>
