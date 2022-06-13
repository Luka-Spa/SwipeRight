import http from 'k6/http';
import { group, sleep, check } from 'k6';
import { Counter } from 'k6/metrics';

// A simple counter for http requests
export const requests = new Counter('http_reqs');
const port = __ENV.PORT || 8089
const ip = __ENV.HOST || "localhost"
const token = __ENV.TOKEN || "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IlRlc3QgVXNlciIsImlhdCI6MTUxNjIzOTAyMiwicGVybWlzc2lvbnMiOlsiY3JlYXRlOnVzZXJzIiwicmVhZDp1c2VycyJdfQ.jRY6RmVeZXwh_gu8sM-kVlcnTlVlOcXJV57c6bCNHnY"


const u = () => {
  return {
  email: "example.test@gmail.com",
  profile_image_url: "https://profile_image_url.com/test.jpg",
  first_name: "John",
  last_name: "Doe",
  gender: "m",
  bio: "Test bio",
  location: {
    city: "Eindhoven",
    lat: 51.441643,
    long: 5.469722
  },
  anthem: "September",
  school: "Fontys",
  preferences: {
    visible: true,
    age_min: 19,
    age_max: 22,
    gender_preference: "f",
    max_distance: 12
  }}
}
// you can specify stages of your test (ramp up/down patterns) through the options object
// target is the number of VUs you are aiming for
export const options = {
  stages: [
    { target: 100, duration: '1m' },
    { target: 15, duration: '1m' },
    { target: 0, duration: '1m' },
  ],
  thresholds: {
    http_req_failed: ['rate<0.01'],
    http_req_duration: ['p(95)<800'],
  },
};

export default function () {
  // our HTTP request, note that we are saving the response to res, which can be accessed later
  group('Performance Testing', function () {
    group('Get user', function () {
      const res = http.get(`http://${ip}:${port}/api/user`, {headers: {Authorization: "Bearer " + token }});
      check(res, {
        'is status code 200': (r) => r.status === 200,
      });
    });
    sleep(1)
    // group('Post user', function () {
    //   const res = http.post(`http://${ip}:${port}/api/user`,u , {headers: {Authorization: "Bearer " + token }});
    //   check(res, {
    //     'is status code 201': (r) => r.status === 201,
    //   });
    // });
  });
}
