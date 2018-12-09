#!/usr/bin/perl

use strict;
use warnings;

my $input_file = 'input.txt';

open(PUZZLE_INPUT, "<./$input_file")
    or die("Cannot open file!\n");

my $input_line = <PUZZLE_INPUT>;
close(PUZZLE_INPUT);

chomp $input_line;

my $result = 0;

my @digits = split //, $input_line;
my $a_size = @digits;
my $shift = $a_size/2;
my $counter = 0;

foreach my $digit (@digits)
{
    if($counter < $shift)
    {
        if($digit == $digits[$counter + $shift])
        {
            $result += $digit;
        }
    }
    else
    {
        if($digit == $digits[$counter - $shift])
        {
            $result += $digit;
        }
    }

    $counter++;
}

print 'Result: ' . $result . "\n";
