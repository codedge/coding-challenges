<?php

declare(strict_types=1);

namespace Tests;

use Exception;
use PHPUnit\Framework\TestCase;

final class T6Test extends TestCase
{
    public function testFibonacci(): void
    {
        $this->assertSame(0, fibonacci(0));
        $this->assertSame(1, fibonacci(1));
        $this->assertSame(2, fibonacci(3));
        $this->assertSame(21, fibonacci(8));
        $this->assertSame(88169, fibonacci(38));
    }

    public function testThrowException(): void
    {
        $this->expectException(Exception::class);
        fibonacci(-1);
    }
}
