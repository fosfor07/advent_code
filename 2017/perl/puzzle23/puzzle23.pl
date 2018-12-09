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


# Part 1
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
    elsif($inst eq 'sub')
    {
        $registers{$arg1} = $registers{$arg1} - $arg2;
    }
    elsif($inst eq 'mul')
    {
        $registers{$arg1} = $registers{$arg1} * $arg2;
        $played++;
    }
    elsif($inst eq 'jnz')
    {
        if($arg1 =~ /^[a-zA-Z]/)
        {
            if($registers{$arg1} != 0)
            {
                $i = $i + $arg2;
                $i_updated = 1;
            }
        }
        else
        {
            if($arg1 != 0)
            {
                $i = $i + $arg2;
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
        $i++;
    }

    if($i > $#instructions)
    {
        last;
    }
}


# Part 2
my $cnt = 0;
my $x = 0;

for($x=107900; $x<=(107900 + 17000); $x=$x+17)
{
    if( (($x % 2) == 0) || (($x % 3) == 0) )
    {
        $cnt++;
    }
    else
    {
        my $i = 5;

        while($i * $i <= $x)
        {
            if( (($x % $i) == 0) || (($x % ($i + 2)) == 0) )
            {
                $cnt++;
                last;
            }

            $i = $i + 6;
        }
    }
}

print 'Result 1: '. $played . "\n";
print 'Result 2: '. $cnt . "\n";

