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
my @input = split /\n/, $input_line;
my $cond = 0;
my $oper = 0;
my %registers = ();
my $cmp_res = 0;
my $calc_res = 0;
my $result = 0;

foreach my $instr (@input)
{
    ($oper, $cond) = split /if/, $instr;
    $oper =~ s/^\s+|\s+$//g;
    $cond =~ s/^\s+|\s+$//g;

    my @oper_inc = split / /, $oper;
    my @cond_inc = split / /, $cond;

    if(defined $registers{$cond_inc[0]})
    {
        $cmp_res = compare_fun($registers{$cond_inc[0]}, $cond_inc[1], $cond_inc[2]);

        if($cmp_res != 0)
        {
            if(defined $registers{$oper_inc[0]})
            {
                $calc_res = calc_fun($registers{$oper_inc[0]}, $oper_inc[1], $oper_inc[2]);
            }
            else
            {
                $calc_res = calc_fun(0, $oper_inc[1], $oper_inc[2]);
            }

            $registers{$oper_inc[0]} = $calc_res;
        }
    }
    else
    {
        $cmp_res = compare_fun(0, $cond_inc[1], $cond_inc[2]);

        if($cmp_res != 0)
        {
            if(defined $registers{$oper_inc[0]})
            {
                $calc_res = calc_fun($registers{$oper_inc[0]}, $oper_inc[1], $oper_inc[2]);
            }
            else
            {
                $calc_res = calc_fun(0, $oper_inc[1], $oper_inc[2]);
            }

            $registers{$oper_inc[0]} = $calc_res;
        }
    }

    if($calc_res > $result)
    {
        $result = $calc_res;
    }
}

print 'Result: ' . $result . "\n";

sub calc_fun
{
    if($_[1] eq 'inc')
    {
        return $_[0] + $_[2];
    }
    elsif($_[1] eq 'dec')
    {
        return $_[0] - $_[2];
    }
}

sub compare_fun
{
    my $res = 0;

    if($_[1] eq '>')
    {
        if($_[0] > $_[2])
        {
            $res = 1;
        }
        else
        {
            $res = 0;
        }
    }
    elsif($_[1] eq '<')
    {
        if($_[0] < $_[2])
        {
            $res = 1;
        }
        else
        {
            $res = 0;
        }
    }
    elsif($_[1] eq '>=')
    {
        if($_[0] >= $_[2])
        {
            $res = 1;
        }
        else
        {
            $res = 0;
        }
    }
    elsif($_[1] eq '<=')
    {
        if($_[0] <= $_[2])
        {
            $res = 1;
        }
        else
        {
            $res = 0;
        }
    }
    elsif($_[1] eq '==')
    {
        if($_[0] == $_[2])
        {
            $res = 1;
        }
        else
        {
            $res = 0;
        }
    }
    elsif($_[1] eq '!=')
    {
        if($_[0] != $_[2])
        {
            $res = 1;
        }
        else
        {
            $res = 0;
        }
    }
    else
    {
        print "ERROR: Unknown operator!\n";
    }

    return $res;
}

