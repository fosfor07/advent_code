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
my $cancel = 0;
my $inside_grp = 0;
my $inside_gb = 0;
my $result = 0;

foreach my $char (split //, $input_line)
{
    if($cancel != 0)
    {
        $cancel = 0;
        next;
    }

    if($char eq '!')
    {
        $cancel = 1;
    }
    elsif( ($char eq '<') && ($inside_gb == 0) )
    {
        $inside_gb = 1;
    }
    elsif( ($char eq '>') && ($inside_gb != 0) )
    {
        $inside_gb = 0;
    }
    elsif( ($char eq '{') && ($inside_gb == 0) )
    {
        $inside_grp++;
    }
    elsif( ($char eq '}') && ($inside_gb == 0) && ($inside_grp != 0) )
    {
        $inside_grp--;
    }
    else
    {
        if($inside_gb != 0)
        {
            $result++;
        }
    }
}

print 'Result: ' . $result . "\n";

