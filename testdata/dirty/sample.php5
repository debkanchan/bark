<?php
class User {
    public $name;

    function __construct($name) {
        $this->name = $name;
    }

    function greet() {
        return "Hello, " . $this->name;
    }
}

$user = new User("Sample user");


// BARK: This is a sample PHP5 file to test the parser registry.
echo $user->greet();
?>