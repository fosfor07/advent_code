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

my %layers = ();
my $max = 0;
my ($layer, $depth) = 0;

foreach my $line (split /\n/, $input_line)
{
    if($line =~ /:/)
    {
        ($layer, $depth) = split /:/, $line;
        $layer =~ s/^\s+|\s+$//g;
        $depth =~ s/^\s+|\s+$//g;

        $layers{$layer} = $depth;
    }
    else
    {
        print 'ERROR: Line without : found: ' . $line . "\n";
    }

    $max = $layer;
}

my $idx = 0;
my $state = 0;
my $delay = 0;

while(1)
{
    while(1)
    {
        if(defined $layers{$idx})
        {
            if($layers{$idx} == 2)
            {
                if((($idx + $delay) % 2) == 0)
                {
                    $state += ($idx * $layers{$idx});
                }
                else
                {
                    $state += 0;
                }
            }
            elsif( (((($idx + $delay) % ($layers{$idx} - 1)) == 0) &&
                ((($idx + $delay) / ($layers{$idx} - 1)) % 2 == 0)) )
            {
                $state += ($idx * $layers{$idx}) + 1;
            }
            else
            {
                $state += 0;
            }
        }
        else
        {
            $state += 0;
        }

        if($idx == $max)
        {
            last;
        }

        $idx++;
    }

    if($state == 0)
    {
        last;
    }
    else
    {
        $delay++;
        $idx = 0;
        $state = 0;
    }
}

print 'Result: ' . $delay . "\n";
