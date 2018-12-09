#!/usr/bin/perl

use strict;
use warnings;

my $steps = 377;
my $buf_size = 50000000;
my $updt_steps = 3;
my $idx = 0;
my $after_zero = 0;
my $all_elems = 0;

# we start with 0 in the buffer
$all_elems++;

for(my $i = 1;$i<$buf_size;$i++)
{
    $updt_steps = $steps;

    while($updt_steps > ($all_elems - 1))
    {
        $updt_steps -= $all_elems;
    }

    $idx += $updt_steps;

    if($idx > ($all_elems - 1))
    {
        $idx -= $all_elems;
    }

    if($idx == ($all_elems - 1))
    {
        $idx = 0;
        $all_elems++;
        $after_zero = $i;
    }
    else
    {
        $idx++;
        $all_elems++;
    }
}

print 'Result: ' . $after_zero . "\n";

