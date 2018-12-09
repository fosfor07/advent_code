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
my @cur_states = split /\t/, $input_line;

my @prev_states;
my $max_bank_idx = 15;
my $highest = 0;
my $tmp_highest = 0;
my $highest_idx = 0;
my $idx = 0;
my $prev_idx = 0;
my $idx_prev_found = 0;
my $num_prev = 0;
my $match_found = 0;
my $result = 0;



foreach my $state (@cur_states)
{
    $prev_states[0][$idx] = $state;

    if($idx == 0)
    {
        $highest = $state;
        $highest_idx = 0;
    }
    else
    {
        if($state > $highest)
        {
            $highest = $state;
            $highest_idx = $idx;
        }
    }
    $idx++;
}

$num_prev++;

while(1)
{
    $cur_states[$highest_idx] = 0;
    if(($highest_idx + 1) > $max_bank_idx)
    {
        $idx = 0;
    }
    else
    {
        $idx = $highest_idx + 1;
    }

    $tmp_highest = 0;
    $highest_idx = $max_bank_idx + 1;

    for(my $i=0;$i<=$max_bank_idx;$i++)
    {
        if($idx > $max_bank_idx)
        {
            $idx = 0;
        }

        if($highest > 0)
        {
            $cur_states[$idx] += 1;
            $highest--;
        }

        if($cur_states[$idx] > $tmp_highest)
        {
            $tmp_highest = $cur_states[$idx];
            $highest_idx = $idx;
        }
        if(($cur_states[$idx] == $tmp_highest) && ($highest_idx > $idx))
        {
            $highest_idx = $idx;
        }

        $idx++;
    }

    $highest = $tmp_highest;
    $result++;

    my $the_same = 1;

    $prev_idx = 0;
    foreach my $prev_bank (@prev_states)
    {
        $idx = 0;
        foreach my $prev_state (@$prev_bank)
        {
            if($prev_state != $cur_states[$idx])
            {
                $the_same = 0;
                last;
            }
            $idx++;
        }

        $prev_idx++;

        if($the_same == 0)
        {
            $the_same = 1;
            next;
        }
        else
        {
            $idx_prev_found = $prev_idx - 1;
            $match_found = 1;
            last;
        }
    }

    if($match_found != 0)
    {
        last;
    }

    $idx = 0;
    foreach my $state (@cur_states)
    {
        $prev_states[$num_prev][$idx] = $state;
        $idx++;
    }
    $num_prev++;
}

print 'Result part 1: ' . $result . "\n";
print 'Result part 2: ' . ($result - $idx_prev_found) . "\n";

