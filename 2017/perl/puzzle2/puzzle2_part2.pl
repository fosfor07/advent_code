#!/usr/bin/perl

use strict;
use warnings;

my $input_file = 'input.txt';

open(PUZZLE_INPUT, "<./$input_file")
    or die("Cannot open file!\n");

my $counter_n = 0;
my $counter_d = 0;
my $result = 0;
my $higest = 0;
my $lowest = 0;
my $division_result = 0;
my @numbers;

while (my $input_line = <PUZZLE_INPUT>)
{
    chomp $input_line;

    @numbers = split /\t/, $input_line;
    $division_result = 0;

    foreach my $number (@numbers)
    {
        foreach my $divider (@numbers)
        {
            # We don't want to divide element by itself. It will always return 1.
            if($counter_n != $counter_d)
            {
                if($number >= $divider)
                {
                    if(($number % $divider) == 0)
                    {
                        $division_result = $number / $divider;
                    }
                }
            }

            $counter_d++;
        }

        $counter_n++;
        $counter_d = 0;
    }
    $counter_n = 0;

    $result += $division_result;
}
close(PUZZLE_INPUT);

print 'Result: ' . $result . "\n";
