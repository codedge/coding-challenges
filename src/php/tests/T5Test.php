<?php

declare(strict_types=1);

namespace Tests;

use PHPUnit\Framework\TestCase;

final class T5Test extends TestCase
{
    public function testKcomplementPositiveNumber(): void
    {
        $this->assertSame(7, k_complement(6, [1, 8, -3, 0, 1, 3, -2, 4, 5]));
    }

    public function testKcomplementNegativeNumber(): void
    {
        $this->assertSame(2, k_complement(-1, [-1, 2, 0]));
    }
}
