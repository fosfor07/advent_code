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


my @fields;
my $x = 200;
my $y = 200;
my $dir = 'u';
my $infections = 0;

# for debugging
my $min_x = 200;
my $min_y = 200;
my $max_x = 224;
my $max_y = 224;


foreach my $row (split /\n/, $input)
{
    chomp $row;
    $x = 200;

    foreach my $field (split //, $row)
    {
        $fields[$x][$y] = $field;
        $x++;
    }
    $y++;
}

# We start in the middle.
$x = 212;
$y = 212;

for(my $i=0; $i<10000; $i++)
{
    if( !defined $fields[$x][$y] )
    {
        $fields[$x][$y] = '.';
    }

    $dir = change_dir($fields[$x][$y], $dir);

    if($fields[$x][$y] eq '#')
    {
        $fields[$x][$y] = '.';
    }
    else
    {
        $fields[$x][$y] = '#';
        $infections++;
    }

    ($x, $y) = move($x, $y, $dir);


    if(($x < 0) || ($y < 0))
    {
        print "ERROR: Reached end of the board. Board is too small!\n";
        last;
    }

    # for debugging
    if($x > $max_x)
    {
        $max_x = $x;
    }
    if($x < $min_x)
    {
        $min_x = $x;
    }
    if($y > $max_y)
    {
        $max_y = $y;
    }
    if($y < $min_y)
    {
        $min_y = $y;
    }
}

print 'Infections: ' . $infections . "\n";


sub move
{
    my $x = shift;
    my $y = shift;
    my $dir = shift;

    if($dir eq 'u')
    {
        $y--;
    }
    elsif($dir eq 'd')
    {
        $y++;
    }
    elsif($dir eq 'l')
    {
        $x--;
    }
    elsif($dir eq 'r')
    {
        $x++;
    }

    return ($x, $y);
}


sub change_dir
{
    my $field = shift;
    my $dir = shift;

    if($field eq '#')
    {
        if($dir eq 'u')
        {
            $dir = 'r';
        }
        elsif($dir eq 'r')
        {
            $dir = 'd';
        }
        elsif($dir eq 'd')
        {
            $dir = 'l';
        }
        elsif($dir eq 'l')
        {
            $dir = 'u';
        }
    }
    else
    {
        if($dir eq 'u')
        {
            $dir = 'l';
        }
        elsif($dir eq 'r')
        {
            $dir = 'u';
        }
        elsif($dir eq 'd')
        {
            $dir = 'r';
        }
        elsif($dir eq 'l')
        {
            $dir = 'd';
        }
    }

    return $dir;
}

