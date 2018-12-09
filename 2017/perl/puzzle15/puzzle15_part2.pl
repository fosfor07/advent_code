#!/usr/bin/perl

use strict;
use warnings;

my $result = 0;

my $gen_a = 516;
my $gen_b = 190;

my $factor_a = 16807;
my $factor_b = 48271;
my $m = 2147483647;
my $gen_val_a = $gen_a;
my $gen_val_b = $gen_b;

for(my $j=0;$j<5000000;$j++)
{
    while(1)
    {
        $gen_val_a = (($gen_val_a % $m) * $factor_a) % $m;
        if(($gen_val_a % 4) == 0)
        {
            last;
        }
    }

    while(1)
    {
        $gen_val_b = (($gen_val_b % $m) * $factor_b) % $m;
        if(($gen_val_b % 8) == 0)
        {
            last;
        }
    }

    #my $val_a_hex = sprintf("%04x", $gen_val_a);
    #my $val_b_hex = sprintf("%04x", $gen_val_b);

    #if(substr($val_a_hex, -4, 4) eq substr($val_b_hex, -4, 4))
    if(($gen_val_a & 65535) == ($gen_val_b & 65535))
    {
        $result++;
    }
}

print 'Result: ' . $result . "\n";

