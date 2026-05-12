# Zadanie 1, dodatkowe 2

```bash
docker buildx create --name labbuilder --driver docker-container --use
```
<img width="751" height="92" alt="image" src="https://github.com/user-attachments/assets/1e655b36-b53d-4f3e-86fe-d756ae9d3b54" />

```bash
docker buildx inspect --bootstrap
```
<img width="936" height="1046" alt="image" src="https://github.com/user-attachments/assets/08b01998-eda4-427f-a784-915aa1422387" />

```bash
docker buildx build \
  --platform linux/amd64,linux/arm64 \
  -t kedraantoni/weather-app:latest \
  --push \
  --cache-from type=registry,ref=kedraantoni/weather-app:buildcache \
  --cache-to type=registry,ref=kedraantoni/weather-app:buildcache,mode=max \
  --build-arg BUILDKIT_INLINE_CACHE=1 \
  .
```

```bash
docker buildx imagetools inspect antek03/pawcho:latest
```
<img width="1314" height="634" alt="image" src="https://github.com/user-attachments/assets/3501db7f-25fe-47ba-8fa0-afcfa165fc39" />
<img width="933" height="1000" alt="image" src="https://github.com/user-attachments/assets/a0bf86fc-b91d-4286-95dd-5c123e99cd85" />
