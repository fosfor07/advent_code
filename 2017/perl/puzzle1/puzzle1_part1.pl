#!/usr/bin/perl

use strict;
use warnings;

my $input_file = 'input.txt';

open(PUZZLE_INPUT, "<./$input_file")
    or die("Cannot open file!\n");

my $input_line = <PUZZLE_INPUT>;
close(PUZZLE_INPUT);

chomp $input_line;

my $counter = 0;
my $first_digit = 0;
my $prev_digit = 0;
my $cur_digit = 0;
my $result = 0;

foreach (split //, $input_line)
{
    if($counter == 0)
    {
        $first_digit = $_;
        $prev_digit = $_;
        $counter++;
    }
    else
    {
        $cur_digit = $_;

        if($cur_digit == $prev_digit)
        {
            $result += $prev_digit;
        }

        $prev_digit = $cur_digit;
    }
}

if($cur_digit == $first_digit)
{
    $result += $first_digit;
}

print 'Result: ' . $result . "\n";
