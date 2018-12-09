#!/usr/bin/perl

use strict;
use warnings;

my $input_file = 'input.txt';

open(PUZZLE_INPUT, "<./$input_file")
    or die("Cannot open file!\n");

my $counter_wrd = 0;
my $counter_cmp = 0;
my $pass_ok = 1;
my $result = 0;

while (my $input_line = <PUZZLE_INPUT>)
{
    chomp $input_line;
    my @words = split / /, $input_line;

OUTER_FOREACH:
    foreach my $word (@words)
    {
        chomp $word;
        foreach my $cmp_word (@words)
        {
            chomp $cmp_word;

            if($counter_wrd != $counter_cmp)
            {
                if($word eq $cmp_word)
                {
                    $pass_ok = 0;
                    $counter_cmp = 0;
                    last OUTER_FOREACH;
                }
            }

            $counter_cmp++;
        }

        $counter_wrd++;
        $counter_cmp = 0;
    }

    $counter_wrd = 0;

    if($pass_ok != 0)
    {
        $result++;
    }
    else
    {
        $pass_ok = 1;
    }
}
close(PUZZLE_INPUT);

print 'Result: ' . $result . "\n";
