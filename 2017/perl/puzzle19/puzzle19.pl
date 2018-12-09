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

chomp $input;

my $result = '';
my $steps = 1;
my @fields;
my $x = 0;
my $y = 0;
my $start_x = 0;

foreach my $line (split /\n/, $input)
{
    $x = 0;
    foreach my $char (split //, $line)
    {
        $fields[$x][$y] = $char;
        $x++;
    }
    $y++;
}

my $max_x = $x;
my $max_y = $y;

for(my $i=0;$i<$x;$i++)
{
    if($fields[$i][0] ne '|')
    {
        next;
    }
    else
    {
        $start_x = $i;
    }
}

my $down = 1;
my $right = 1;
my $vertical = 1;

$x = $start_x;
$y = 1;

while(1)
{
    if( (!defined $fields[$x][$y]) ||
        ($fields[$x][$y] eq ' ') )
    {
        last;
    }
    elsif($fields[$x][$y] eq '|')
    {
        if($vertical != 0)
        {
            if($down != 0)
            {
                $y++;
            }
            else
            {
                $y--;
            }
        }
        else
        {
            if($right != 0)
            {
                $x++;
            }
            else
            {
                $x--;
            }
        }
    }
    elsif($fields[$x][$y] eq '-')
    {
        if($vertical == 0)
        {
            if($right != 0)
            {
                $x++;
            }
            else
            {
                $x--;
            }
        }
        else
        {
            if($down != 0)
            {
                $y++;
            }
            else
            {
                $y--;
            }
        }
    }
    elsif($fields[$x][$y] eq '+')
    {
        # down
        if(($vertical != 0) && ($down != 0))
        {
            # down
            if( (($y+1) < $max_y) && ($fields[$x][$y+1] ne ' ') )
            {
                $y++;
            }
            # right
            elsif( (($x+1) < $max_x) && ($fields[$x+1][$y] ne ' ') )
            {
                $x++;
                $vertical = 0;
                $right = 1;
            }
            # left
            elsif( (($x-1) >= 0) && ($fields[$x-1][$y] ne ' ') )
            {
                $x--;
                $vertical = 0;
                $right = 0;
            }
            else
            {
                last;
            }
        }
        # up
        elsif(($vertical != 0) && ($down == 0))
        {
            # up
            if( (($y-1) >= 0) && ($fields[$x][$y-1] ne ' ') )
            {
                $y--;
            }
            # right
            elsif( (($x+1) < $max_x) && ($fields[$x+1][$y] ne ' ') )
            {
                $x++;
                $vertical = 0;
                $right = 1;
            }
            # left
            elsif( (($x-1) >= 0) && ($fields[$x-1][$y] ne ' ') )
            {
                $x--;
                $vertical = 0;
                $right = 0;
            }
            else
            {
                last;
            }
        }
        # right
        elsif(($vertical == 0) && ($right != 0))
        {
            # right
            if( (($x+1) < $max_x) && ($fields[$x+1][$y] ne ' ') )
            {
                $x++;
            }
            # up
            elsif( (($y+1) < $max_y) && ($fields[$x][$y+1] ne ' ') )
            {
                $y++;
                $vertical = 1;
                $down = 1;
            }
            # down
            elsif( (($y-1) >= 0) && ($fields[$x][$y-1] ne ' ') )
            {
                $y--;
                $vertical = 1;
                $down = 0;
            }
            else
            {
                last;
            }
        }
        # left
        elsif(($vertical == 0) && ($right == 0))
        {
            # left
            if( (($x-1) >= 0) && ($fields[$x-1][$y] ne ' ') )
            {
                $x--;
            }
            # down
            elsif( (($y+1) < $max_y) && ($fields[$x][$y+1] ne ' ') )
            {
                $y++;
                $vertical = 1;
                $down = 1;
            }
            # up
            elsif( (($y-1) >= 0) && ($fields[$x][$y-1] ne ' ') )
            {
                $y--;
                $vertical = 1;
                $down = 0;
            }
            else
            {
                last;
            }
        }
    }
    else
    {
        $result .= $fields[$x][$y];

        # down
        if(($vertical != 0) && ($down != 0))
        {
            $y++;
        }
        # up
        elsif(($vertical != 0) && ($down == 0))
        {
            $y--;
        }
        # right
        elsif(($vertical == 0) && ($right != 0))
        {
            $x++;
        }
        # left
        elsif(($vertical == 0) && ($right == 0))
        {
            $x--;
        }
    }

    $steps++;
}

print 'Result: ' . $result . "\n";
print 'Steps: ' . $steps . "\n";

