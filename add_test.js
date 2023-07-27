import http from 'k6/http';
import { sleep } from 'k6';

export default function () {
    const url = 'http://127.0.0.1:42000/add';
    const payload = JSON.stringify({
      FirstInt: 1,
      SecondInt: 2,
    });
  
    const params = {
      headers: {
        'Content-Type': 'application/json',
      },
    };
  
    http.post(url, payload, params);
    sleep(2)
  }