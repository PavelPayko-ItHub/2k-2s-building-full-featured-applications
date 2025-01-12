const api = 'http://localhost:8090'
const target_userlogin = '/user/login'
const target_usercreate = '/user/create'
const target_orderplace = '/order/place'
const target_categories = '/categories' //'/api/category/'
const target_category = '/category/list/' //'/api/category/'
//const target_subcategory = '/api/subcategory/'
const target_products = '/product/get/' //'/api/products/'

function buildPath(target) {
    return api.concat('/', target)
}

function makeStruct(keys) {
    if (!keys) return null;
    const k = keys.split(', ');
    const count = k.length;

    /** @constructor */
    function constructor() {
        for (let i = 0; i < count; i++) this[k[i]] = arguments[i];
    }
    return constructor;
}

async function userLogin(credentials) {
    const remote = api.concat('', `${target_userlogin}?username=${credentials.email}&password=${credentials.password}`)
    console.log(remote)
    const requestOptions = {
        method: "POST",
        headers: {
            'Content-type': 'application/json',
        },
    };
    const response = await fetch(remote, requestOptions)
    return response.json();
}

async function userCreate(credentials) {
    const remote = api.concat('', `${target_usercreate}?username=${credentials.email}&password=${credentials.password}`)
    console.log(remote)
    const requestOptions = {
        method: "POST",
        headers: {
            'Content-type': 'application/json',
        },
    };
    const response = await fetch(remote, requestOptions)
    return response.json();
}

async function placeOrder(jsonbody, accessToken) {
    const remote = api.concat('', target_orderplace)
    console.log(remote)
    const requestOptions = {
        method: "POST",
        headers: {
            'Content-type': 'application/json',
            "Authorization": `Bearer ${accessToken}`
        },
        body: jsonbody,
    };
    const response = await fetch(remote, requestOptions)
    console.log(response)
    return response.json();
}

async function getCategories() {
    const remote = api.concat('', target_categories)
    console.log(remote)
    const response = await fetch(remote)
    return response.json();
}

async function getCategory(id) {
    const remote = api.concat('', `${target_category}${id}`)
    console.log(remote)
    const response = await fetch(remote)
    return response.json();
}

async function getProduct(id) {
    const response = await fetch(api.concat('', `${target_products}${id}`))
    return response.json();
}

export {buildPath, makeStruct, userLogin, userCreate, placeOrder, getCategories, getCategory, getProduct};
