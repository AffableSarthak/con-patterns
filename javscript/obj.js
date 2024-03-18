const obj = {
    value: 42,
    getValue: () => this.value
};

obj.value = 99;
console.log(obj.getValue()); // This will return 42
