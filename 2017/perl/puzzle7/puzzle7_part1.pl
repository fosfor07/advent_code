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
my %all_towers = ();
my @lower_lvl;

foreach my $line (split /\n/, $input_line)
{
    if($line =~ /->/)
    {
        my ($part_1, $part_2) = split /->/, $line;

        my @line_parts = split / /, $part_1;
        $line_parts[1] =~ /\((\d+)\)/;
        my $weight = $1;
        $all_towers{ $line_parts[0] } = $weight;

        foreach my $low_tower (split /,/, $part_2)
        {
            $low_tower =~ s/^\s+|\s+$//g;
            push @lower_lvl, $low_tower;
        }
    }
    else
    {
        my @line_parts = split / /, $line;
        $line_parts[1] =~ /\((\d+)\)/;
        my $weight = $1;
        $all_towers{ $line_parts[0] } = $weight;
    }
}

foreach my $tower ( keys %all_towers )
{
    my $found = 0;

    foreach my $low_lvl (@lower_lvl)
    {
        if($tower eq $low_lvl)
        {
            $found = 1;
            last;
        }
    }

    if($found == 0)
    {
        print 'Root: ' . $tower . "\n";
    }
}

