<?php

declare(strict_types=1);

function number_to_ordinal(int $number): string
{
    if (0 === $number) {
        return '0';
    }

    $ending = 'th';

    if (!in_array(($number % 100), [11, 12, 13])) {
        switch ($number % 10) {
            // Handle 1st, 2nd, 3rd
            case 1:
                $ending = 'st';
                break;

            case 2:
                $ending = 'nd';
                break;

            case 3:
                $ending = 'rd';
                break;
        }
    }

    return $number.$ending;
}
