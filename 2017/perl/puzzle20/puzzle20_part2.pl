#!/usr/bin/perl

use strict;
use warnings;
use List::Util qw( min max );

my $input_file = 'input.txt';

local $/ = undef;
open(PUZZLE_INPUT, "<./$input_file")
    or die("Cannot open file!\n");

my $input = <PUZZLE_INPUT>;
close(PUZZLE_INPUT);
local $/ = "\n";

chomp $input;


my @pos;
my @vel;
my @acc;
my %duplicates;

foreach my $line (split /\n/, $input)
{
    my ($p, $v, $a) = split / /, $line;
    chomp $p;
    chomp $v;
    chomp $a;

    my ($param, $val) = split /=/, $p;
    $val =~ s/^<|>,$|>$//g;
    my ($x, $y, $z) = split /,/, $val;
    push @pos, $val;

    if( defined $duplicates{$val} )
    {
        $duplicates{$val} += 1;
    }
    else
    {
        $duplicates{$val} = 1;
    }

    ($param, $val) = split /=/, $v;
    $val =~ s/^<|>,$|>$//g;
    ($x, $y, $z) = split /,/, $val;
    push @vel, $val;

    ($param, $val) = split /=/, $a;
    $val =~ s/^<|>,$|>$//g;
    ($x, $y, $z) = split /,/, $val;
    push @acc, $val;
}

foreach my $key (keys %duplicates)
{
    if($duplicates{$key} > 1)
    {
        while($duplicates{$key} > 0)
        {
            remove_element($key, \@pos, \@vel, \@acc);
            $duplicates{$key} -= 1;
        }
    }
}


my $cnt = 0;

while(1)
{
    my ($xp, $yp, $zp);
    my ($xv, $yv, $zv);
    my ($xa, $ya, $za);

    %duplicates = ();

    for(my $j=0; $j<=$#pos; $j++)
    {
        ($xp, $yp, $zp) = split /,/, $pos[$j];
        ($xv, $yv, $zv) = split /,/, $vel[$j];
        ($xa, $ya, $za) = split /,/, $acc[$j];

        $xv += $xa;
        $yv += $ya;
        $zv += $za;
        $vel[$j] = join(',',$xv,$yv,$zv);

        $xp += $xv;
        $yp += $yv;
        $zp += $zv;
        $pos[$j] = join(',',$xp,$yp,$zp);

        if( defined $duplicates{$pos[$j]} )
        {
            $duplicates{$pos[$j]} += 1;
        }
        else
        {
            $duplicates{$pos[$j]} = 1;
        }
    }

    foreach my $key (keys %duplicates)
    {
        if($duplicates{$key} > 1)
        {
            while($duplicates{$key} > 0)
            {
                remove_element($key, \@pos, \@vel, \@acc);
                $duplicates{$key} -= 1;
            }
        }
    }

    $cnt++;

    if($cnt > 500)
    {
        print 'Number of particles: ' . scalar @pos . "\n";
        last;
    }
}


sub remove_element
{
    my $elem = shift;
    my $pos = shift;
    my $vel = shift;
    my $acc = shift;

    my $idx = 0;

    foreach my $position (@$pos)
    {
        if($elem eq $position)
        {
            splice @$pos, $idx, 1;
            splice @$vel, $idx, 1;
            splice @$acc, $idx, 1;
            last;
        }
        $idx++;
    }
}

