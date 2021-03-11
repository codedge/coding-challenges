<?php

declare(strict_types=1);

namespace Tests;

use PHPUnit\Framework\TestCase;

final class T4Test extends TestCase
{
    public function testCanCheckJumpAtEndOfArray(): void
    {
        $this->assertSame(4, jump_out_of_array([2, 3, -1, 1, 6, 4]));
    }

    public function testCanCheckJumpAtBeginningOfArray(): void
    {
        $this->assertSame(1, jump_out_of_array([-1, 3, -1, 1, 6, 4]));
    }

    public function testCannotJumpOutOfArray(): void
    {
        $this->assertSame(-1, jump_out_of_array([1, 2, 1, 1, 1, 0]));
    }
}
