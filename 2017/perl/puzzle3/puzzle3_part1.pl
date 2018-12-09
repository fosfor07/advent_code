#!/usr/bin/perl

use strict;
use warnings;

my $input = 347991;
my $square_side = 1;
my $square_fields = 1;
my $corner = 0;
my $middle = 0;
my $result = 0;

while ($square_fields < $input)
{
    $square_side += 2;
    $square_fields = $square_side * $square_side;
}

$corner = $square_fields;

while ($corner > $input)
{
    $corner -= ($square_side - 1);
}

$middle = $corner + (($square_side - 1) / 2);

# distance from the middle to center + distance from the middle to our field
my $dist_to_m = 0;
if($middle > $input)
{
    $dist_to_m = $middle - $input;
}
else
{
    $dist_to_m = $input - $middle;
}
$result = (($square_side - 1) / 2) + $dist_to_m;

print 'Result: ' . $result . "\n";
