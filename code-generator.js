const rand = () => {
	return Math.floor(Math.random() * 10);
};

for (let i = 0; i < 200; i++) {
	let code = '';
	for (let i = 0; i < 10; i++) {
		let r = rand();
		while (i === 0 && r === 0) r = rand();
		code += r.toString();
	}
	console.log(code);
}

