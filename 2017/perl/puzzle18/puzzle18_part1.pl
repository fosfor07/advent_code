#!/usr/bin/perl

use strict;
use warnings;

my $input_file = 'input.txt';

local $/ = undef;
open(PUZZLE_INPUT, "<./$input_file")
    or die("Cannot open file!\n");

my $input = <PUZZLE_INPUT>;
close(PUZZLE_INPUT);
local $/ = "\n";

chomp $input;

my %registers;
my @instructions;
my $played = 0;
my $i = 0;

foreach my $line (split /\n/, $input)
{
    chomp $line;
    push @instructions, $line;
}

while(1)
{
    my $i_updated = 0;
    my ($inst, $arg1, $arg2) = split / /, $instructions[$i];

    if( (!defined $registers{$arg1}) &&
       (defined $arg1) && ($arg1 =~ /^[a-zA-Z]/) )
    {
        $registers{$arg1} = 0;
    }

    if((defined $arg2) && ($arg2 =~ /^[a-zA-Z]/))
    {
        if(!defined $registers{$arg2})
        {
            $registers{$arg2} = 0;
        }
        $arg2 = $registers{$arg2};
    }

    if($inst eq 'set')
    {
        $registers{$arg1} = $arg2;
    }
    elsif($inst eq 'add')
    {
        $registers{$arg1} = $registers{$arg1} + $arg2;
    }
    elsif($inst eq 'mul')
    {
        $registers{$arg1} = $registers{$arg1} * $arg2;
    }
    elsif($inst eq 'mod')
    {
        $registers{$arg1} = $registers{$arg1} % $arg2;
    }
    elsif($inst eq 'snd')
    {
        $played = $registers{$arg1};
    }
    elsif($inst eq 'rcv')
    {
        if($registers{$arg1} != 0)
        {
            last;
        }
    }
    elsif($inst eq 'jgz')
    {
        if($registers{$arg1} > 0)
        {
            $i = $i + $arg2;
            $i_updated = 1;
        }
    }
    else
    {
        print "ERROR: Unknown instruction!\n";
    }

    if($i_updated == 0)
    {
        $i++;
    }
}

print 'Result: '. $played . "\n";

