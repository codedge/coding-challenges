<?php

declare(strict_types=1);

function mirror(array $numbers): array {
    $max = max(array_keys($numbers));

    for($i=0; $i<=(floor(count($numbers)/2))-1; $i++) {
        $tmp = $numbers[$max-$i];
        $numbers[$max-$i] = $numbers[$i];
        $numbers[$i] = $tmp;
    }

    return $numbers;
}
