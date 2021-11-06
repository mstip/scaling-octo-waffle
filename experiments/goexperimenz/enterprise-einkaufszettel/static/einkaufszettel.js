function toggleDone(index) {
    const params = new URLSearchParams();
    params.append('itemIndex', index);
    axios.post('/toggle-item-done', params);
    setTimeout(() => window.location.href = '/', 50);
}

function deleteItem(index) {
    const params = new URLSearchParams();
    params.append('itemIndex', index);
    axios.post('/delete-item', params);
    setTimeout(() => window.location.href = '/', 50);
}