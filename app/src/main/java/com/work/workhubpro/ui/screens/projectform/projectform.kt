package com.work.workhubpro.ui.screens.projectform



import android.widget.Toast
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.verticalScroll
import androidx.compose.material3.Button
import androidx.compose.material3.Icon
import androidx.compose.material3.OutlinedTextField
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Brush
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.painter.Painter
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import androidx.hilt.navigation.compose.hiltViewModel
import androidx.navigation.NavController
import com.work.workhubpro.R
import com.work.workhubpro.SharedViewModel
import com.work.workhubpro.ui.composables.HeadingTextComposable
import com.work.workhubpro.ui.composables.NormalTextComposable

@Composable
fun CreateProject(
    navController: NavController,
    sharedViewModel: SharedViewModel
) {

    var name by remember { mutableStateOf("") }
    var description by remember { mutableStateOf("") }
    var projectlead by remember { mutableStateOf("") }
    val createProjectViewModel:ProjectFormmViewMoel = hiltViewModel()
    var response = createProjectViewModel.response.collectAsState().value
    val scrollState = rememberScrollState()
    val context = LocalContext.current


    val gradient = Brush.verticalGradient(
        colors = listOf(
            Color(0xFF00B0F0), // Start color
            Color(0xFF0077C2)  // End color
        )
    )

    Scaffold(
        modifier = Modifier
            .fillMaxSize()
            .background(Color.White)
            .padding(30.dp),
    ) {
        Column(
            modifier = Modifier
                .fillMaxSize()
                .padding(it)
                .verticalScroll(scrollState),

            ) {
            NormalTextComposable(value = stringResource(id = R.string.Create_New_project))
            HeadingTextComposable(value = stringResource(id = R.string.create_project))

            MyTextField(
                labelValue = stringResource(id = R.string.project_name),
                painterResource(id = R.drawable.company_symbol),
                textValue = name,
                onValueChange = { name = it }
            )
            MyTextField(
                labelValue = stringResource(id = R.string.project_description),
                painterResource(id = R.drawable.outline_edit_black_24dp),
                textValue = description,
                onValueChange = { description = it }
            )
            MyTextField(
                labelValue = stringResource(id = R.string.project_description),
                painterResource(id = R.drawable.outline_edit_black_24dp),
                textValue = projectlead,
                onValueChange = { projectlead = it }
            )


            Button(
                modifier = Modifier
                    .fillMaxWidth()
                    .padding(15.dp)
                    .background(
                        color =
                        Color.hsl(248f, 0.95f, 0.60f), // Valid form color

                        shape = RoundedCornerShape(10.dp)
                    ),
                shape = RoundedCornerShape(10.dp),
                onClick = {
                    println("ashwathama")
                    println(sharedViewModel.user.value)
                    createProjectViewModel.createProject(name, description, projectlead,sharedViewModel.user.value?.id!!)
                },
            ) {
                Text(text = "Create Project",color=Color.White)
            }

        }

    }
}

@Composable
fun MyTextField(
    labelValue: String,
    painterResource: Painter,
    textValue: String,
    onValueChange: (String) -> Unit,
    isError: Boolean = false
) {
    OutlinedTextField(
        value = textValue,
        onValueChange = onValueChange,
        modifier = Modifier
            .fillMaxSize()
            .padding(10.dp)
        ,
        leadingIcon = {
            Icon(painter = painterResource, contentDescription = null, modifier = Modifier.size(24.dp))
        },
        label = {
            Text(text = labelValue)
        },
        isError = isError
    )
}



