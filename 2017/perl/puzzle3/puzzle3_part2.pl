#!/usr/bin/perl

use strict;
use warnings;

my $input = 347991;
my @fields;
my $new_field = 1;
my $x = 100;
my $y = 100;

$fields[$x][$y] = $new_field;
$x++;

while ($new_field < $input)
{
    $new_field = 0;

    # x-1
    if(defined $fields[$x-1][$y+1])
    {
        $new_field += $fields[$x-1][$y+1];
    }
    if(defined $fields[$x-1][$y])
    {
        $new_field += $fields[$x-1][$y];
    }
    if(defined $fields[$x-1][$y-1])
    {
        $new_field += $fields[$x-1][$y-1];
    }

    # x+1
    if(defined $fields[$x+1][$y+1])
    {
        $new_field += $fields[$x+1][$y+1];
    }
    if(defined $fields[$x+1][$y])
    {
        $new_field += $fields[$x+1][$y];
    }
    if(defined $fields[$x+1][$y-1])
    {
        $new_field += $fields[$x+1][$y-1];
    }

    # x
    if(defined $fields[$x][$y+1])
    {
        $new_field += $fields[$x][$y+1];
    }
    if(defined $fields[$x][$y-1])
    {
        $new_field += $fields[$x][$y-1];
    }

    $fields[$x][$y] = $new_field;


    if(defined $fields[$x-1][$y])
    {
        if(defined $fields[$x][$y+1])
        {
            $x++;
        }
        else
        {
            $y++;
        }
    }
    else
    {
        if(defined $fields[$x][$y-1])
        {
            $x--;
        }
        else
        {
            if(defined $fields[$x+1][$y])
            {
                $y--;
            }
            else
            {
                $x++;
            }
        }
    }
}

print 'Result: ' . $new_field . "\n";
