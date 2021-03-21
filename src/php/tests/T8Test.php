<?php

declare(strict_types=1);

namespace Tests;

use Exception;
use PHPUnit\Framework\TestCase;

final class T8Test extends TestCase
{
    public function testEvenInputAmount(): void
    {
        $input = [8,5,1,4];
        $this->assertSame(array_reverse($input), mirror($input));
    }

    public function testOddInputAmount(): void
    {
        $input = [6,2,7,9,3];
        $this->assertSame(array_reverse($input), mirror($input));
    }
}
