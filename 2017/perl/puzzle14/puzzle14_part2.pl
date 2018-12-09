#!/usr/bin/perl

use strict;
use warnings;

## disable perl's warning mechanism
no warnings 'recursion';

my $input = 'amgozmfv-';

my $last_idx = 255;
my ($cnt, $idx) = (0, 0);
my $skip_size = 0;
my @list = ();
my @removed = ();
my @lenghts = ();
my $result = 0;

my @grid;


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
        $binary .= sprintf("%08b", $hash_elem);
    }

    my $j = 0;
    foreach my $char (split //, $binary)
    {
        if($char eq 1)
        {
            $result++;
        }

        $grid[$j][$k] = $char;
        $j++;
    }
}

print 'Result: ' . $result . "\n";

my $regions = 0;
for(my $y=0;$y<128;$y++)
{
    for(my $x=0;$x<128;$x++)
    {
        if($grid[$x][$y] != 0)
        {
            traverse($x, $y, @grid);
            $regions++;
        }
    }
}

print 'Regions: ' . $regions . "\n";

#for(my $k=0;$k<128;$k++)
#{
#    for(my $j=0;$j<128;$j++)
#    {
#        print $grid[$j][$k];
#    }
#    print "\n";
#}


sub traverse
{
    my $root_x = $_[0];
    my $root_y = $_[1];
    my $grid = $_[2];

    $grid[$root_x][$root_y] = 0;

    if( (defined $grid[$root_x - 1][$root_y]) &&
        ($grid[$root_x - 1][$root_y] != 0) && ($root_x != 0) )
    {
        traverse($root_x - 1, $root_y, $grid);
    }

    if( (defined $grid[$root_x + 1][$root_y]) &&
        ($grid[$root_x + 1][$root_y] != 0) && ($root_x != 127) )
    {
        traverse($root_x + 1, $root_y, $grid);
    }

    if( (defined $grid[$root_x][$root_y - 1]) &&
        ($grid[$root_x][$root_y - 1] != 0) && ($root_y != 0) )
    {
        traverse($root_x, $root_y - 1, $grid);
    }

    if( (defined $grid[$root_x][$root_y + 1]) &&
        ($grid[$root_x][$root_y + 1] != 0) && ($root_y != 127) )
    {
        traverse($root_x, $root_y + 1, $grid);
    }
}

