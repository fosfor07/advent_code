#!/usr/bin/perl

use strict;
use warnings;

my $steps = 377;
my $buf_size = 2018;
my $updt_steps = 3;
my $idx = 0;
my @buffer = ();

push @buffer, 0;

for(my $i = 1;$i<$buf_size;$i++)
{
    $updt_steps = $steps;

    while($updt_steps > $#buffer)
    {
        $updt_steps -= ($#buffer + 1);
    }

    $idx += $updt_steps;

    if($idx > $#buffer)
    {
        $idx -= ($#buffer + 1);
    }

    if($idx == $#buffer)
    {
        unshift @buffer, $i;
        $idx = 0;
    }
    else
    {
        $idx++;
        splice @buffer, $idx, 0, $i;
    }
}

if($idx == $#buffer)
{
    $idx = 0;
}
else
{
    $idx++;
}

print 'Result: ' . $buffer[$idx] . "\n";

