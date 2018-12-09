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

my @programs;

for(my $i=0;$i<16;$i++)
{
    $programs[$i] = chr(ord('a') + $i);
}

foreach my $move (split /,/, $input_line)
{
    my $cmd = substr($move, 0, 1);

    if($cmd eq 's')
    {
        spin(\@programs, substr($move, 1));
    }
    elsif($cmd eq 'x')
    {
        $move =~ m/(\d+)\//;
        my $a_idx = $1;
        $move =~ m/\/(\d+)/;
        my $b_idx = $1;

        exchange(\@programs, $a_idx, $b_idx);
    }
    elsif($cmd eq 'p')
    {
        partner(\@programs, substr($move, 1, 1), substr($move, 3, 1));
    }
}



sub spin
{
    my $programs = shift;
    my $spin_num = shift;

    my @removed = splice @$programs, -$spin_num , $spin_num;
    unshift @$programs, @removed;
}

sub exchange
{
    my $programs = shift;
    my $a = shift;
    my $b = shift;

    my $a_val = @{$programs}[$a];
    my $b_val = @{$programs}[$b];

    @{$programs}[$a] = $b_val;
    @{$programs}[$b] = $a_val;
}

sub partner
{
    my $programs = shift;
    my $a_val = shift;
    my $b_val = shift;
    my ($a, $b);

    for(my $i=0;$i<16;$i++)
    {
        if($a_val eq @{$programs}[$i])
        {
            $a = $i;
        }

        if($b_val eq @{$programs}[$i])
        {
            $b = $i;
        }
    }

    @{$programs}[$a] = $b_val;
    @{$programs}[$b] = $a_val;
}

foreach my $elem (@programs)
{
    print $elem;
}

