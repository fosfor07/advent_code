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


my @ports;
my $cnt = 0;

foreach my $line (split /\n/, $input)
{
    chomp $line;
    ($ports[$cnt][0], $ports[$cnt][1]) = (split /\//, $line);
    $cnt++;
}

my @strengths;
my @lengths;
my $max_strength = 0;
my $max_length = 0;
my $l_strength = 0;
my @bridge = [(0, 0)];


build_bridge(0, \@strengths, \@lengths, \@bridge, \@ports);

$cnt = 0;
foreach my $strength (@strengths)
{
    if($strength > $max_strength)
    {
        $max_strength = $strength;
    }

    if($lengths[$cnt] >= $max_length)
    {
        $max_length = $lengths[$cnt];
        if($strength > $l_strength)
        {
            $l_strength = $strength;
        }
    }

    $cnt++;
}

print 'Result 1: ' . $max_strength . "\n";
print 'Result 2: ' . $l_strength . ' ' . $max_length . "\n";


sub build_bridge
{
    my $port = shift;
    my $strengths = shift;
    my $lengths = shift;
    my $bridge = shift;
    my @nodes = @{$_[0]};

    my @opts = options($port, \@nodes);

    if(@nodes && @opts)
    {
        foreach my $opt (@opts)
        {
            # Add new element to the bridge.
            push @$bridge, [ $opt->[0], $opt->[1] ];

            my $strength = bridge_strength(\@$bridge);
            push @$strengths, bridge_strength(\@$bridge);
            push @$lengths, ((scalar @$bridge) - 1);

            # Remove this option and continue building the bridge.
            remove_element($opt->[0], $opt->[1], \@nodes);
            build_bridge($opt->[2], \@$strengths, \@$lengths, \@$bridge, \@nodes);

            # Add element again, so it can be used for next option.
            push @nodes, [ $opt->[0], $opt->[1] ];
            # Restore original bridge, so it can be used for next option.
            pop @$bridge;
        }
    }
}


sub options
{
    my $port = shift;
    my @nodes = @{$_[0]};

    my @opts;
    my ($left, $right, $free_port);

    my $cnt = 0;

    while(1)
    {
        ($left, $right, $free_port) = find_next($port, \@nodes);

        if(($left == -1) || ($cnt > 100))
        {
            last;
        }
        else
        {
            push @opts, [$left, $right, $free_port];
        }

        $cnt++;
    }

    return @opts;
}


sub find_next
{
    my $port = shift;
    my $nodes = shift;

    my $cnt = 0;

    foreach my $node (@$nodes)
    {
        if($node->[0] == $port)
        {
            splice @$nodes, $cnt, 1;

            # [left port, right port, free port]
            return ($node->[0], $node->[1], $node->[1]);
        }
        elsif($node->[1] == $port)
        {
            splice @$nodes, $cnt, 1;

            # [left port, right port, free port]
            return ($node->[0], $node->[1], $node->[0]);
        }

        $cnt++;
    }

    return (-1, -1, -1);
}


sub bridge_strength
{
    my $bridge = shift;

    my $strength = 0;

    foreach my $element (@$bridge)
    {
        $strength += $element->[0] + $element->[1];
    }

    return $strength;
}


sub remove_element
{
    my $left = shift;
    my $right = shift;
    my $nodes = shift;

    my $cnt = 0;

    foreach my $node (@$nodes)
    {
        if(($node->[0] == $left) && ($node->[1] == $right))
        {
            splice @$nodes, $cnt, 1;
        }
        $cnt++;
    }
}

