#!/usr/bin/perl

use strict;
use warnings;


my $input_file = 'input.txt';

local $/ = undef;
open(PUZZLE_INPUT, "<./$input_file")
    or die("Cannot open file!\n");

my $input_line = <PUZZLE_INPUT>;
close(PUZZLE_INPUT);
local $/ = "\n";

chomp $input_line;

my $x = 0;
my $y = 0;
my $z = 0;
my $result_part1 = 0;
my $result_part2 = 0;

foreach my $step (split /,/, $input_line)
{
    if($step eq 'n')
    {
        $x++;
        $y--;
    }
    elsif($step eq 'ne')
    {
        $x++;
        $z--;
    }
    elsif($step eq 's')
    {
        $y++;
        $x--;
    }
    elsif($step eq 'se')
    {
        $y++;
        $z--;
    }
    elsif($step eq 'nw')
    {
        $z++;
        $y--;
    }
    elsif($step eq 'sw')
    {
        $z++;
        $x--;
    }

    if(abs($x) > $result_part2)
    {
        $result_part2 = abs($x);
    }

    if(abs($y) > $result_part2)
    {
        $result_part2 = abs($y);
    }

    if(abs($z) > $result_part2)
    {
        $result_part2 = abs($z);
    }
}

if(abs($x) > $result_part1)
{
    $result_part1 = abs($x);
}

if(abs($y) > $result_part1)
{
    $result_part1 = abs($y);
}

if(abs($z) > $result_part1)
{
    $result_part1 = abs($z);
}

print 'Result part 1: ' . $result_part1 . "\n";
print 'Result part 2: ' . $result_part2 . "\n";
