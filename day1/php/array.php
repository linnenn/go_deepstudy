<?php

$arr = [
    "id" => " ",
    "code" => "12",
    "timestamp" => "202045101245",
    "sign" => "b832291d0bb3c0cd9cf824be671cdb18"
];

$arr = array_filter($arr,function($arg){
    if (ctype_space($arg) || empty($arg)){
        return false;
    }else{
        return true;
    }
});

var_dump($arr);