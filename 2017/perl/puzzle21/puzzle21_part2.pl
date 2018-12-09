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


my @grid;
$grid[0][0] = '.';
$grid[1][0] = '#';
$grid[2][0] = '.';
$grid[0][1] = '.';
$grid[1][1] = '.';
$grid[2][1] = '#';
$grid[0][2] = '#';
$grid[1][2] = '#';
$grid[2][2] = '#';

my @new_grid;
my @square = @grid;
my $grid_size = 3;
my $sq_size;
my $out_size;

my $result = 0;

my @in_rule;
my @out_rule;


for(my $iter=0; $iter<18; $iter++)
{
    @new_grid = ();

    # set the square size based on grid size
    if(($grid_size % 2) == 0)
    {
        $sq_size = 2;
    }
    else
    {
        $sq_size = 3;
    }

    # number of squares
    my $num_sqs = ($grid_size/$sq_size) * ($grid_size/$sq_size);

    my $s_x = 0;
    my $s_y = 0;
    my $shift_x = 0;
    my $shift_y = 0;

    # convert each square on the grid
    for(my $s=0; $s<$num_sqs; $s++)
    {
        # read square from the grid
        @square = ();
        my $on_num = 0;
        my $on_x = 0;
        my $on_y = 0;
        for(my $y=0; $y<$sq_size; $y++)
        {
            for(my $x=0; $x<$sq_size; $x++)
            {
                $square[$x][$y] = $grid[$x+$s_x][$y+$s_y];
                if(($sq_size == 2) && ($square[$x][$y] eq '#'))
                {
                    $on_num++;
                    $on_x = $x;
                    $on_y = $y;
                }
            }
        }

        my $fnd = 0;
        # we have 2 rules when on_num = 2, we need to resolve it
        if(($sq_size == 2) && ($on_num == 2))
        {
            for(my $y=0; $y<$sq_size; $y++)
            {
                for(my $x=0; $x<$sq_size; $x++)
                {
                    if($square[$x][$y] eq '#')
                    {
                        if( ($x != $on_x) && ($y != $on_y) )
                        {
                            $on_num = 5;
                        }

                        $fnd = 1;
                        last;
                    }
                }

                if($fnd == 1)
                {
                    last;
                }
            }
        }


        # check rules
        if($sq_size == 3)
        {
            # apply rules to the square
            foreach my $line (split /\n/, $input)
            {
                @in_rule = ();
                @out_rule = ();

                my ($in, $out) = split /=>/, $line;
                $in =~ s/^\s+|\s+$//g;
                $out =~ s/^\s+|\s+$//g;

                # store in rule
                my ($i, $j) = (0, 0);

                foreach my $in_r (split '/', $in)
                {
                    $i = 0;

                    foreach my $char (split '', $in_r)
                    {
                        $in_rule[$i][$j] = $char;
                        $i++;
                    }
                    $j++;
                }

                # rule is skipped if sizes of rule and square are different
                if($sq_size != $i)
                {
                    next;
                }

                # store out rule
                $j = 0;
                foreach my $out_r (split '/', $out)
                {
                    $i = 0;

                    foreach my $char (split '', $out_r)
                    {
                        $out_rule[$i][$j] = $char;
                        $i++;
                    }
                    $j++;
                }
                $out_size = $i;

                # check if in rule matches current square
                my $match_found = check_rule(\@in_rule, \@square, $sq_size);

                if($match_found == 1)
                {
                    for(my $y=0; $y<$out_size; $y++)
                    {
                        for(my $x=0; $x<$out_size; $x++)
                        {
                            $new_grid[$s_x+$x+$shift_x][$s_y+$y+$shift_y] = $out_rule[$x][$y];
                        }
                    }
                    last;
                }
            }
        }
        # pick one of 6 available
        else
        {
            $out_size = 3;

            # ../.. => ##./##./.##
            if($on_num == 0)
            {
                $out_rule[0][0] = '#';
                $out_rule[1][0] = '#';
                $out_rule[2][0] = '.';
                $out_rule[0][1] = '#';
                $out_rule[1][1] = '#';
                $out_rule[2][1] = '.';
                $out_rule[0][2] = '.';
                $out_rule[1][2] = '#';
                $out_rule[2][2] = '#';
            }
            # #./.. => .../.#./##.
            elsif($on_num == 1)
            {
                $out_rule[0][0] = '.';
                $out_rule[1][0] = '.';
                $out_rule[2][0] = '.';
                $out_rule[0][1] = '.';
                $out_rule[1][1] = '#';
                $out_rule[2][1] = '.';
                $out_rule[0][2] = '#';
                $out_rule[1][2] = '#';
                $out_rule[2][2] = '.';
            }
            # ##/#. => .##/#.#/#..
            elsif($on_num == 3)
            {
                $out_rule[0][0] = '.';
                $out_rule[1][0] = '#';
                $out_rule[2][0] = '#';
                $out_rule[0][1] = '#';
                $out_rule[1][1] = '.';
                $out_rule[2][1] = '#';
                $out_rule[0][2] = '#';
                $out_rule[1][2] = '.';
                $out_rule[2][2] = '.';
            }
            # ##/## => ..#/.#./.##
            elsif($on_num == 4)
            {
                $out_rule[0][0] = '.';
                $out_rule[1][0] = '.';
                $out_rule[2][0] = '#';
                $out_rule[0][1] = '.';
                $out_rule[1][1] = '#';
                $out_rule[2][1] = '.';
                $out_rule[0][2] = '.';
                $out_rule[1][2] = '#';
                $out_rule[2][2] = '#';
            }
            # ##/.. => .../.##/#.#
            elsif($on_num == 2)
            {
                $out_rule[0][0] = '.';
                $out_rule[1][0] = '.';
                $out_rule[2][0] = '.';
                $out_rule[0][1] = '.';
                $out_rule[1][1] = '#';
                $out_rule[2][1] = '#';
                $out_rule[0][2] = '#';
                $out_rule[1][2] = '.';
                $out_rule[2][2] = '#';
            }
            # #/#. => ##./#../#..
            elsif($on_num == 5)
            {
                $out_rule[0][0] = '#';
                $out_rule[1][0] = '#';
                $out_rule[2][0] = '.';
                $out_rule[0][1] = '#';
                $out_rule[1][1] = '.';
                $out_rule[2][1] = '.';
                $out_rule[0][2] = '#';
                $out_rule[1][2] = '.';
                $out_rule[2][2] = '.';
            }

            for(my $y=0; $y<$out_size; $y++)
            {
                for(my $x=0; $x<$out_size; $x++)
                {
                    $new_grid[$s_x+$x+$shift_x][$s_y+$y+$shift_y] = $out_rule[$x][$y];
                }
            }
        }

        $s_x += $sq_size;
        $shift_x++;
        if($s_x >= $grid_size)
        {
            $s_x = 0;
            $shift_x = 0;
            $s_y += $sq_size;
            $shift_y++;
        }
    }

    $grid_size = ($grid_size / $sq_size) * $out_size;
    $result = 0;
    @grid = ();

    print 'Iteration: ' . $iter . "\n";
    print 'grid size: ' . $grid_size  . "\n";
    for(my $y=0; $y<$grid_size; $y++)
    {
        for(my $x=0; $x<$grid_size; $x++)
        {
            $grid[$x][$y] = $new_grid[$x][$y];
            #print $grid[$x][$y];
            if($grid[$x][$y] eq '#')
            {
                $result++;
            }
        }
        #print "\n";
    }
    #print "\n";
}

print 'Result: ' . $result . "\n";



sub check_rule
{
    my $rule = shift;
    my $square = shift;
    my $size = shift;

    my $match_found = 0;
    my @mod_rule;

    $match_found = cmp_matrices($rule, $square, $size);

    if($match_found == 0)
    {
        @mod_rule = flip_matrix_v($rule, $size);
        $match_found = cmp_matrices(\@mod_rule, $square, $size);
    }

    if($match_found == 0)
    {
        @mod_rule = flip_matrix_h($rule, $size);
        $match_found = cmp_matrices(\@mod_rule, $square, $size);
    }

    if($match_found == 0)
    {
        # rotate rule and check again
        for(my $i=0; $i<3; $i++)
        {
            rotate_matrix($rule, $size);
            $match_found = cmp_matrices($rule, $square, $size);
            if($match_found == 1)
            {
                last;
            }

            @mod_rule = flip_matrix_v($rule, $size);
            $match_found = cmp_matrices(\@mod_rule, $square, $size);
            if($match_found == 1)
            {
                last;
            }

            @mod_rule = flip_matrix_h($rule, $size);
            $match_found = cmp_matrices(\@mod_rule, $square, $size);
            if($match_found == 1)
            {
                last;
            }
        }
    }

    return $match_found;
}


sub rotate_matrix
{
    my $src_m = shift;
    my $size = shift;

    my $tmp;

    for(my $i=0; $i<$size/2; $i++)
    {
        for(my $j=$i; $j<($size - $i - 1); $j++)
        {
            $tmp                                = $src_m->[$i][$j];
            $src_m->[$i][$j]                    = $src_m->[$j][$size-$i-1];
            $src_m->[$j][$size-$i-1]            = $src_m->[$size-$i-1][$size-$j-1];
            $src_m->[$size-$i-1][$size-$j-1]    = $src_m->[$size-$j-1][$i];
            $src_m->[$size-$j-1][$i]            = $tmp;
        }
    }
}


sub flip_matrix_v
{
    my $src_m = shift;
    my $size = shift;

    # copy rule so we can flip it
    my @matrix;
    for(my $j=0;$j<$size;$j++)
    {
        for(my $i=0;$i<$size;$i++)
        {
            $matrix[$i][$j] = $src_m->[$i][$j];
        }
    }

    my $tmp;

    for(my $j=0; $j<$size; $j++)
    {
        for(my $i=0; $i<$size/2; $i++)
        {
            $tmp = $matrix[$i][$j];
            $matrix[$i][$j] = $matrix[$size-1-$i][$j];
            $matrix[$size-1-$i][$j] = $tmp;
        }
    }

    return @matrix;
}


sub flip_matrix_h
{
    my $src_m = shift;
    my $size = shift;

    # copy rule so we can flip it
    my @matrix;
    for(my $j=0;$j<$size;$j++)
    {
        for(my $i=0;$i<$size;$i++)
        {
            $matrix[$i][$j] = $src_m->[$i][$j];
        }
    }

    my $tmp;

    for(my $i=0; $i<$size; $i++)
    {
        for(my $j=0; $j<$size/2; $j++)
        {
            $tmp = $matrix[$i][$j];
            $matrix[$i][$j] = $matrix[$i][$size-1-$j];
            $matrix[$i][$size-1-$j] = $tmp;
        }
    }

    return @matrix;
}


sub cmp_matrices
{
    my $rule = shift;
    my $square = shift;
    my $size = shift;

    my $match_found = 1;

    for(my $y=0; $y<$size; $y++)
    {
        for(my $x=0; $x<$size; $x++)
        {
            if($square->[$x][$y] ne $rule->[$x][$y])
            {
                $match_found = 0;
                last;
            }
        }

        if($match_found == 0)
        {
            last;
        }
    }

    return $match_found;
}

