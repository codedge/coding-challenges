<?php

declare(strict_types=1);

function calculate(string $expression)
{
    $result = 0;
    $splittedExpressions = explode(' ', $expression);

    if (empty($expression)) {
        return $result;
    }

    // Last item is not an operand but a numeric value, then directly output it.
    // Special handling to check if float or int value.
    if (is_numeric(end($splittedExpressions))) {
        return false !== strpos(end($splittedExpressions), '.') ? (float) end($splittedExpressions) : (int) end($splittedExpressions);
    }

    $numbers = [];
    $validOperations = ['+', '-', '*', '/'];

    foreach ($splittedExpressions as $single) {
        if (false === is_numeric($single) && !in_array($single, $validOperations)) {
            return false;
        }

        if (is_numeric($single)) {
            $numbers[] = $single;
        } else {
            $number1 = array_pop($numbers);
            $number2 = array_pop($numbers);

            switch ($single) {
                case '+':
                    $result = $number2 + $number1;
                    break;

                case '-':
                    $result = $number2 - $number1;
                    break;

                case '*':
                    $result = $number2 * $number1;
                    break;

                case '/':
                    $result = $number2 / $number1;
                    break;
            }

            $numbers[] = $result;
        }
    }

    return $result;
}
