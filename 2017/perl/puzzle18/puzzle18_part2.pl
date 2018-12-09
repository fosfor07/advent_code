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

my %registers = ();
my @instructions = ();
my $sent = 0;
my $i0 = 0;
my $i1 = 0;
my $executes = 0;

$registers{'p'}[0] = 0;
$registers{'p'}[1] = 1;

my @queue_0 = ();
my @queue_1 = ();

foreach my $line (split /\n/, $input)
{
    chomp $line;
    push @instructions, $line;
}

while(1)
{
    my $i_updated = 0;
    my ($inst, $arg1, $arg2);

    if($executes == 0)
    {
        ($inst, $arg1, $arg2) = split / /, $instructions[$i0];
    }
    else
    {
        ($inst, $arg1, $arg2) = split / /, $instructions[$i1];
    }

    if( (!defined $registers{$arg1}[$executes]) &&
        (defined $arg1) && ($arg1 =~ /^[a-zA-Z]/) )
    {
        $registers{$arg1}[$executes] = 0;
    }

    if((defined $arg2) && ($arg2 =~ /^[a-zA-Z]/))
    {
        if(!defined $registers{$arg2}[$executes])
        {
            $registers{$arg2}[$executes] = 0;
        }
        $arg2 = $registers{$arg2}[$executes];
    }

    if($inst eq 'set')
    {
        $registers{$arg1}[$executes] = $arg2;
    }
    elsif($inst eq 'add')
    {
        $registers{$arg1}[$executes] = $registers{$arg1}[$executes] + $arg2;
    }
    elsif($inst eq 'mul')
    {
        $registers{$arg1}[$executes] = $registers{$arg1}[$executes] * $arg2;
    }
    elsif($inst eq 'mod')
    {
        $registers{$arg1}[$executes] = $registers{$arg1}[$executes] % $arg2;
    }
    elsif($inst eq 'snd')
    {
        if($executes == 0)
        {
            push @queue_1, $registers{$arg1}[$executes];
        }
        else
        {
            push @queue_0, $registers{$arg1}[$executes];
            $sent++;
        }
    }
    elsif($inst eq 'rcv')
    {
        if($executes == 0)
        {
            if(@queue_0)
            {
                $registers{$arg1}[$executes] = shift @queue_0;
                $i0++;
            }
            else
            {
                $executes = 1;
            }
        }
        else
        {
            if(@queue_1)
            {
                $registers{$arg1}[$executes] = shift @queue_1;
                $i1++;
            }
            else
            {
                $executes = 0;
            }
        }
        $i_updated = 1;

        if((@queue_0 == 0) && (@queue_1 == 0))
        {
            last;
        }
    }
    elsif($inst eq 'jgz')
    {
        if($arg1 =~ /^[a-zA-Z]/)
        {
            if($registers{$arg1}[$executes] > 0)
            {
                if($executes == 0)
                {
                    $i0 = $i0 + $arg2;
                }
                else
                {
                    $i1 = $i1 + $arg2;
                }
                $i_updated = 1;
            }
        }
        else
        {
            if($arg1 > 0)
            {
                if($executes == 0)
                {
                    $i0 = $i0 + $arg2;
                }
                else
                {
                    $i1 = $i1 + $arg2;
                }
                $i_updated = 1;
            }
        }
    }
    else
    {
        print "ERROR: Unknown instruction!\n";
    }

    if($i_updated == 0)
    {
        if($executes == 0)
        {
            $i0++;
        }
        else
        {
            $i1++;
        }
    }

    if( ($i0 > $#instructions) || ($i1 > $#instructions) )
    {
        last;
    }
}

print 'Result: '. $sent . "\n";

