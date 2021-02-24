<?php

declare(strict_types=1);

namespace Tests;

use PHPUnit\Framework\TestCase;

final class T3Test extends TestCase
{
    public function testAddition()
    {
        $this->assertSame(5, calculate('2 3 +'));
    }

    public function testSubstration()
    {
        $this->assertSame(5, calculate('6 1 -'));
    }

    public function testMultiplication()
    {
        $this->assertSame(16, calculate('2 8 *'));
    }

    public function testDivision()
    {
        $this->assertSame(2, calculate('10 5 /'));
    }

    public function testReturnNumber()
    {
        $this->assertSame(3.5, calculate('1 2 3.5'));
        $this->assertSame(4, calculate('2 3 4'));
    }

    public function testWrongOperation()
    {
        $this->assertSame(false, calculate('10 5 %'));
    }

    public function test0()
    {
        $this->assertSame(0, calculate(''));
    }
}
