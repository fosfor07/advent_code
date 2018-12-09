#!/usr/bin/perl

use strict;
use warnings;

my %states;
$states{'A'} = [1, 1, 'B', 0, -1, 'C'];
$states{'B'} = [1, -1, 'A', 1, 1, 'C'];
$states{'C'} = [1, 1, 'A', 0, -1, 'D'];
$states{'D'} = [1, -1, 'E', 1, -1, 'C'];
$states{'E'} = [1, 1, 'F', 1, 1, 'A'];
$states{'F'} = [1, 1, 'A', 1, 1, 'E'];

my @tape;
push @tape, 0;

my $idx = 0;
my $max_idx = 0;
my $state = 'A';
my $offset = 0;


for(my $i=0; $i<12261543; $i++)
{
    if($idx < 0)
    {
        unshift @tape, 0;
        $max_idx++;
        $idx++;
    }
    elsif($idx > $max_idx)
    {
        push @tape, 0;
        $max_idx++;
    }


    if($tape[$idx] == 0)
    {
        $offset = 0;
    }
    elsif($tape[$idx] == 1)
    {
        $offset = 3;
    }

    $tape[$idx] = $states{$state}[$offset];
    $idx += $states{$state}[$offset + 1];
    $state = $states{$state}[$offset + 2];
}

my $cnt = 0;
foreach my $value (@tape)
{
    if($value == 1)
    {
        $cnt++;
    }
}

print 'Result: ' . $cnt . "\n";

