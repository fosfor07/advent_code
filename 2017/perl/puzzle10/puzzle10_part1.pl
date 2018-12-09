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
my $last_idx = 255;
my $idx = 0;
my $skip_size = 0;
my @list = (0..$last_idx);
my @removed;
my @lenghts;


foreach my $length (split /,/, $input_line)
{
    $lenghts[$idx] = $length;
    $idx++;
}

$idx = 0;

foreach my $length (@lenghts)
{
    my ($f_len, $s_len) = (0,0);

    if(($idx + $length) > ($last_idx - 1))
    {
        $f_len = $last_idx - $idx + 1;
        $s_len = $length - $f_len;

        @removed = splice @list, $idx, $f_len;
        push @removed, splice @list, 0, $s_len;

        @removed = reverse @removed;
        push @list, @removed[0..($f_len - 1)];
        unshift @list, @removed[$f_len..$#removed];
    }
    else
    {
        @removed = splice @list, $idx, $length;
        splice @list, $idx, 0, reverse @removed;
    }

    $idx = ($idx + $length + $skip_size) % ($last_idx + 1);

    $skip_size++;
    if($skip_size == ($last_idx + 1))
    {
        $skip_size = 0;
    }
}

print 'Result: ' . ($list[0] * $list[1]) . "\n";

