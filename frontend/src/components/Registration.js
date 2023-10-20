import React, { useState } from 'react';

const Registration = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [dateOfBirth, setDateOfBirth] = useState('');

    const handleRegistration = () => {
        // Implement registration logic here
    };

    return (
        <div>
            <h1>Registration</h1>
            <input type="email" placeholder="Email" value={email} onChange={(e) => setEmail(e.target.value)} />
            <input type="password" placeholder="Password" value={password} onChange={(e) => setPassword(e.target.value)} />
            <input type="date" placeholder="Date of Birth" value={dateOfBirth} onChange={(e) => setDateOfBirth(e.target.value)} />
            <button onClick={handleRegistration}>Register</button>
        </div>
    );
};

export default Registration;
