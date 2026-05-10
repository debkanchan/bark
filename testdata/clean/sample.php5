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


// Normal comment
echo $user->greet();
?>