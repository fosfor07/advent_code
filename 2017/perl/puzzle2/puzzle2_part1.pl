#!/usr/bin/perl

use strict;
use warnings;

my $input_file = 'input.txt';

open(PUZZLE_INPUT, "<./$input_file")
    or die("Cannot open file!\n");

my $counter = 0;
my $result = 0;
my $higest = 0;
my $lowest = 0;
my $difference = 0;

while (my $input_line = <PUZZLE_INPUT>)
{
    chomp $input_line;
    $counter = 0;
    $difference = 0;

    foreach my $number (split /\t/, $input_line)
    {
        if($counter == 0)
        {
            $higest = $number;
            $lowest = $number;
            $counter++;
        }
        else
        {
            if($number > $higest)
            {
                $higest = $number;
            }
            elsif($number < $lowest)
            {
                $lowest = $number;
            }
        }
    }

    $difference = $higest - $lowest;
    $result += $difference;
}
close(PUZZLE_INPUT);

print 'Result: ' . $result . "\n";
