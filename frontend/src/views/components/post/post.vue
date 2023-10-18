<template>

    <div class="body__warpper">
        <div id="Aside">
            <Aside></Aside>
        </div>
        <div id="Content">
            <Content></Content>
        </div>
    </div>

</template>

<script lang="js" setup>
import Aside from "./components/aside.vue"
import Content from "./components/content.vue"

</script>

<style scoped lang="scss">
.body__warpper{
    display: flex;
    padding: 0;
    overflow: hidden;
    justify-content: center;
    #Aside{
        margin-right: 3vw;
    }
    #Content{
        margin-right: 11vw;
    }
}
</style>