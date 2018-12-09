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
my %all_weights = ();
my %all_lvls = ();
my $root = 'bsfpjtc';
my $result = 0;

foreach my $line (split /\n/, $input_line)
{
    if($line =~ /->/)
    {
        my ($part_1, $part_2) = split /->/, $line;

        my @line_parts = split / /, $part_1;
        $line_parts[1] =~ /\((\d+)\)/;
        my $weight = $1;
        $all_weights{ $line_parts[0] } = $weight;

        my $idx = 0;
        foreach my $low_tower (split /,/, $part_2)
        {
            $low_tower =~ s/^\s+|\s+$//g;
            $all_lvls{ $line_parts[0] }[$idx] = $low_tower;
            $idx++;
        }
    }
    else
    {
        my @line_parts = split / /, $line;
        $line_parts[1] =~ /\((\d+)\)/;
        my $weight = $1;
        $all_weights{ $line_parts[0] } = $weight;
    }
}

while(1)
{
    my %weights = ();
    my $br_weight = 0;
    my $found_new = 0;

    foreach my $level (@{$all_lvls{ $root }})
    {
        $br_weight = traverse($level, %all_lvls, %all_weights);

        if(defined $weights{$br_weight})
        {
            $weights{$br_weight} .= ',' . $level;
        }
        else
        {
            $weights{$br_weight} = $level;
        }
    }

    foreach my $new_lvl (values %weights)
    {
        if($new_lvl !~ /,/)
        {
            $root = $new_lvl;
            $found_new = 1;
        }
    }

    while ( my ($key, $value) = each(%weights) )
    {
        print "$key => $value\n";
    }
    print "\n";

    if($found_new == 0)
    {
        last;
    }
}

print 'root: ' . $root . "\n";

$result = $all_weights{$root};
print 'Result: ' . $result . "\n";


sub traverse
{
    my $root = $_[0];
    my $all_lvls = $_[1];
    my $all_weights = $_[2];
    my $weight = $all_weights{$root};
    my $result = 0;

    foreach my $node (@{$all_lvls{ $root }})
    {
        $result += traverse($node, $all_lvls, $all_weights);
    }

    return $result + $weight;
}

