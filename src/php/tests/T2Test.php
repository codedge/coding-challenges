<?php

declare(strict_types=1);

namespace Tests;

use PHPUnit\Framework\TestCase;

final class T2Test extends TestCase
{
    public function test1st()
    {
        $this->assertSame('1st', number_to_ordinal(1));
    }

    public function test2nd()
    {
        $this->assertSame('2nd', number_to_ordinal(2));
    }

    public function test3rd()
    {
        $this->assertSame('3rd', number_to_ordinal(3));
    }

    public function test4th()
    {
        $this->assertSame('4th', number_to_ordinal(4));
    }
}
