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
my @input = split /\n/, $input_line;
my $cur_idx = 0;
my $cur_jump = 0;
my $result = 0;

while( ($cur_idx >= 0) && ($cur_idx <= $#input) )
{
    $cur_jump = $input[$cur_idx];
    if($input[$cur_idx] < 3)
    {
        $input[$cur_idx]++;
    }
    else
    {
        $input[$cur_idx]--;
    }
    $cur_idx += $cur_jump;

    $result++;
}

print 'Result: ' . $result . "\n";
