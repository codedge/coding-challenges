<?php

declare(strict_types=1);

function k_complement(int $k, array $a): int {
    $amountComplimentary = 0;

    /**
     * Hold amount of occurrence of a single number inside starting array '$a'
     * Structure is:
     *  [
     *      [
     *          key: number
     *          value: amount occurrence in original array '$a'
     *      ],
     *      [
     *          ...
     *      ],
     *  ]
     */
    $occurrencesNumbers = [];

    foreach ($a as $number) {
        if (!array_key_exists($number, $occurrencesNumbers)) {
            $occurrencesNumbers[$number] = 1;
        } else {
            $occurrencesNumbers[$number]++;
        }
    }

    $numbers = array_keys($occurrencesNumbers);

    foreach ($numbers as $number) {
        // Building the complementary pair by subtracting one number from the target result
        // and checking if the result (other pair part) exists
        $result = $k - $number;
        if (in_array($result, $numbers)) {
            $amountComplimentary += $occurrencesNumbers[$number] * $occurrencesNumbers[$result];
        }
    }

    return $amountComplimentary;
}
