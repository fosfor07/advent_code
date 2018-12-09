#!/usr/bin/perl

use strict;
use warnings;

my $input = 'amgozmfv-';

my $last_idx = 255;
my ($cnt, $idx) = (0, 0);
my $skip_size = 0;
my @list = ();
my @removed = ();
my @lenghts = ();
my $result = 0;

for(my $k=0;$k<128;$k++)
{
    @list = (0..$last_idx);
    @lenghts = ();
    my $key = $input . $k;

    $cnt = 0;
    foreach my $char (split //, $key)
    {
        $lenghts[$cnt] = ord($char);
        $cnt++;
    }
    push @lenghts, 17;
    push @lenghts, 31;
    push @lenghts, 73;
    push @lenghts, 47;
    push @lenghts, 23;

    $idx = 0;
    $skip_size = 0;

    for(my $x=0; $x<64; $x++)
    {
        foreach my $length (@lenghts)
        {
            my ($f_len, $s_len) = (0,0);

            if(($idx + $length) > $last_idx)
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
    }

    my @dense_hash = ();
    for(my $y=0;$y<16; $y++)
    {
        $dense_hash[$y] = $list[0+($y*16)] ^ $list[1+($y*16)] ^
                        $list[2+($y*16)] ^ $list[3+($y*16)] ^
                        $list[4+($y*16)] ^ $list[5+($y*16)] ^
                        $list[6+($y*16)] ^ $list[7+($y*16)] ^
                        $list[8+($y*16)] ^ $list[9+($y*16)] ^
                        $list[10+($y*16)] ^ $list[11+($y*16)] ^
                        $list[12+($y*16)] ^ $list[13+($y*16)] ^
                        $list[14+($y*16)] ^ $list[15+($y*16)];
    }


    my $binary;

    foreach my $hash_elem (@dense_hash)
    {
        $binary .= sprintf("%b", $hash_elem);
    }

    foreach my $char (split //, $binary)
    {
        if($char eq 1)
        {
            $result++;
        }
    }
}

print 'Result: ' . $result . "\n";

