<?php

$str1 = '1c0111001f010100061a024b53535009181c';
$str2 = '686974207468652062756c6c277320657965';
$expected = '746865206b696420646f6e277420706c6179';

$result = bin2hex(pack('H*', $str1) ^ pack('H*', $str2));

if ($result == $expected) {
	echo "Success!\n";
} else {
	echo "Failed.\n";
}

echo "Expected: " . $expected . "\n";
echo "Got: " . $result . "\n";
?>
