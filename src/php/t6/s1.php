<?php

declare(strict_types=1);

function fib(int $n): int {
    if ($n < 0) {
        throw new Exception("Only non-negative integers are allowed");
    }

    if($n <= 1) {
        return $n;
    }

    return fibonacci($n-1) + fibonacci($n-2);
}

function fibonacci(int $n): int {
    $result = fib($n);

    $str = (string) $result;
    if (strlen($str) > 6 ) {
        return (int) substr($str, -6);
    }

    return (int) $str;
}
