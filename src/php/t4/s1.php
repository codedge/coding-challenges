<?php

declare(strict_types=1);

function jump_out_of_array(array $a): int {
    $sum = 0;
    $jumps = 1;

    // If first entry in array is negative, then the first jump is out of the array
    if($a[0] < 0) {
        return $jumps;
    }

    for ($i=0; $i<= count($a); $i++) {
        $p = $a[$sum];
        $sum = $sum + $p;

        // Handle jumps out of array at the beginning of the array
        if($sum < 0) {
            return $jumps;
        }

        $jumps++;

        // Handle jumps out of array at the end of the array
        if (($sum + $p) > count($a)
        ) {
            return $jumps;
        }

        if(($i + 1) === count($a)) {
            return -1;
        }
    }
}
