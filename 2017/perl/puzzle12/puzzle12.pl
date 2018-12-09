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

my %all_lvls = ();
my $result_part1 = 0;
my $result_part2 = 0;

foreach my $line (split /\n/, $input_line)
{
    if($line =~ /<->/)
    {
        my ($part_1, $part_2) = split /<->/, $line;
        $part_1 =~ s/^\s+|\s+$//g;

        my $idx = 0;
        foreach my $program (split /,/, $part_2)
        {
            $program =~ s/^\s+|\s+$//g;

            $all_lvls{ $part_1 }[$idx] = $program;
            $idx++;
        }
    }
    else
    {
        print 'Line without <-> found: ' . $line . "\n";
    }
}

my %checked = ();

$result_part1 = traverse('0', %all_lvls, %checked);
$result_part2++;

foreach my $node (keys %all_lvls)
{
    if( !defined $checked{$node} )
    {
        traverse($node, %all_lvls, %checked);
        $result_part2++;
    }
}

print 'Result part 1: ' . $result_part1 . "\n";
print 'Result part 2: ' . $result_part2 . "\n";

sub traverse
{
    my $root = $_[0];
    my $all_lvls = $_[1];
    my $checked = $_[2];
    my $result = 1;

    $checked{$root} = 1;

    foreach my $node (@{$all_lvls{ $root }})
    {
        if( !defined $checked{$node} )
        {
            $result += traverse($node, $all_lvls);
        }
    }

    return $result;
}
