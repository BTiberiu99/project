<template>
  <v-card raised="raised" :key="up" class="">
      <div v-for="(row,index) in config.Matrix" :key="index + 'row'" class="Puzzle__row">
            <component :is="type" v-for="(nr,index2) in  config.Matrix[index]" 
            :key="index * config.Matrix.length + index2" 
            class="Puzzle__number"
            :value="config.Matrix[index][index2]"
            @input=" (input) => validateInput(input,index,index2) "
            dense
            flat>
              {{nr}}
            </component >
      </div>  
      <div v-if="message" style="margin-top:5px;color:#f00;max-width:100px;">
        {{message}}
      </div>
  </v-card>
</template>

<script>
  export default {

    props:{
      config:{
        required:true,
        type:Object
      },
      input:{
        type:Boolean,
        default:false
      }
    },
    data () {

      return {
        up:0,
        message:'',
      }
    },

    created(){
     
    },

    computed:{
      type(){
          return this.input ? 'v-text-field' : 'div'
      },
     
    },
    methods: {
      validate(){
       var missing = []
        
        for (let i=0;i<9;i++){
          if(!this.isNumberInMatrix(i)){
            missing.push(i.toString())
          }
        }
         
        if(missing.length){
          this.message = 'You are missing numbers:' + missing.join(',')
        }else {
          this.message = ''
        }
     
        return missing.length === 0
      },
      matrixToString(){
        let i=0,j=0
        var str = ""
        const m = this.config.Matrix
        for(;i<m.length;i++){
          for (j=0;j<m[i].length;j++){
            str += m[i][j] + " "
          }
          str += "\n"
        }

        return str
      },
      isNumberInMatrix(nr){
        let i=0,j=0
        const m = this.config.Matrix
        for(;i<m.length;i++){
   
          for (j=0;j<m[i].length;j++){

            if (m[i][j] === nr){
              return [i,j]
            }
          }
         
        }

        return false

      },
      valid(input){
            var nr
            if (typeof input === 'number'){
              nr = input
            }else {
              nr = parseInt(input)
            }
             

            if (isNaN(nr)){
              return false
            }

            if (nr < 1 || nr > 8){
              return false
            }
            return true
        },

        log(){
          console.log(arguments)
        },
        validateInput(input,index,index2){    
             if(this.valid(input)){
               input = parseInt(input)
               var items =this.isNumberInMatrix(input)
               if(items){
                 var aux = this.config.Matrix[index][index2]
                 var [i,j] = items
                 this.config.Matrix[index][index2] = this.config.Matrix[i][j]
                  this.config.Matrix[i][j] = aux
               }else {
                 this.config.Matrix[index][index2] = input
               }
               
               
             }
           
             this.validate()
             
             this.up++
        }
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
.Puzzle {
  &__row{
    display: block;
  }

  &__number{
    display: inline-block;
    border:1.2px solid #000;
    width:30px;
    height: 30px;
    text-align: center;

    &--changed{
      color:#f00;
    }

    &--with{
      color:#00f
    }
  }
}
</style>

<style lang="scss" >
.v-text-field__slot{
  text-align: center;
  input{
    text-align: center;
  }
}
</style>