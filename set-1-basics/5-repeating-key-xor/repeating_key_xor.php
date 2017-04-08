<?php

$text = <<<EOT
Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal
EOT;
$key = 'ICE';
$key = str_split($key);

$expected = "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f";

$dupe_key = $key;
$current_key_char = "";
$results = [];
foreach (str_split($text) as $character) {
    $current_key_char = array_shift($dupe_key);

    // Check to see if we need to reset the key array
    if (sizeof($dupe_key) == 0) {
        $dupe_key = $key;
    }


    $results[] = bin2hex($character ^ $current_key_char);
}

$got = join('', $results);

if ($got === $expected) {
    echo "We got it: " . $got . "\n";
} else {
    echo "Expected: " . $expected . "\n";
    echo "Got     : " . $got . "\n";
}

?>
